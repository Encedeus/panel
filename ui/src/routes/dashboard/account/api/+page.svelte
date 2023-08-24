<script lang="ts">
    import Button from "$lib/components/generic/Button.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Input from "$lib/components/generic/Input.svelte";
    import TextArea from "$lib/components/generic/TextArea.svelte";
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
    import PlusIcon from "$lib/components/heroicons/PlusIcon.svelte";
    import ConfirmationModal from "$lib/components/internal/credentials/ConfirmationModal.svelte";
    import KeyTab from "$lib/components/internal/credentials/KeyTab.svelte";
    import { api } from "$lib/services/api";
    import { getSignedInUser } from "$lib/services/auth_service";
    import Toast from "$lib/components/generic/Toast.svelte";
    import { AccountApiKey, HttpError } from "@encedeus/js-api";
    import { isIP } from "is-ip";
    import { onMount } from "svelte";

    let modalOpen = false;

    let keyDescription = "";
    let allowedIpsBox = "";

    let error: string | undefined | null = undefined;
    let descriptionError = false;
    let ipError = false;

    let apiKeys: AccountApiKey[] = [];

    let deleteId = "";

    onMount(async () => {
        const resp = await api.apiKeyService.findAccountApiKeysByUserId((await getSignedInUser()).id)
        if (resp.error) {
            error = "Failed fetching API keys";
        }

        apiKeys = resp.keys!;
        console.log(apiKeys);
    });


    function validateInput(): boolean {
        if (!keyDescription) {
            error = "No description provided";
            descriptionError = true;
            return false;
        }
        if (keyDescription.length > 24) {
            error = "Description is too long";
            descriptionError = true;
            return false;
        }

        for (const ip of allowedIpsBox.trim().split("\n")) {
            if (!isIP(ip) && ip) {
                error = "Invalid IP address"
                ipError = true;
                return false;
            }
        }

        return true;
    }

    async function createKey() {
        if (!validateInput()) {
            return;
        }

        const key = new AccountApiKey();
        key.setDescription(keyDescription);
        key.setAllowedIps(allowedIpsBox.trim().split("\n"));
        key.setUserId((await getSignedInUser()).id);

        const resp = await api.apiKeyService.createAccountApiKey(key);
        if (resp.error) {
            error = resp.error.message;
            return;
        }

        key.setKey(resp.key!);
        apiKeys = [...apiKeys, key];

        keyDescription = "";
        allowedIpsBox = "";
    }

    function clearInput() {
        if (error) {
            error = null;
            descriptionError = false;
            ipError = false;
        }
    }

    async function deleteKey() {
        for (let i = 0; i < apiKeys.length; i++) {
            if (apiKeys[i].id === deleteId) {
                apiKeys = apiKeys.splice(i, 1);
            }
        }

        const resp = await api.apiKeyService.deleteAccountApiKey(deleteId);
        if (resp.error) {
            error = "Failed deleting API key";
        }

        deleteId = "";
    }
</script>

<main class="p-10 flex flex-col items-center">
    <div class="flex min-[1860px]:flex-row flex-col gap-8 items-end justify-end">
        <div class="w-full h-full">
            <CardHeader className="mb-5 self-start" size="lg">
                API Credentials
            </CardHeader>
            <Card className="w-full h-full" height="md" width="lg">
                <span class="text-sm" slot="title">
                    Create API Key
                </span>
                <PlusIcon slot="icon"/>
                <div class="mt-5 w-full h-full flex flex-col items-center justify-between" slot="content">
                    <div class="w-full h-full flex flex-col items-center justify-center gap-4">
                        <Input bind:value={keyDescription} error={descriptionError} label="Description"
                               on:input={clearInput}
                               placeholder="API Key Description" size="xl"/>
                        <TextArea bind:value={allowedIpsBox} className="basis-full resize-none" error={ipError}
                                  label="Allowed IPs"
                                  on:input={clearInput}
                                  placeholder="Leave blank to allow any IP address to use this API key. Otherwise, provide each IP address on a new line."
                                  size="lg"/>
                        <Button className="mt-6 flex justify-center items-center" color="indigo"
                                on:click={async () => await createKey()}
                                size="sm">Create
                        </Button>
                    </div>
                </div>
            </Card>
        </div>
        <div>
            <Card fixedHeight={true} height="md" width="lg">
                <span class="text-sm" slot="title">
                    API Keys
                </span>
                <KeyIcon slot="icon"/>
                <div class="flex flex-col items-center justify-center h-full gap-4 p-8" slot="content">
                    {#each apiKeys as key}
                        <KeyTab className="flex-grow w-full" id={key.id} key={key.key}
                                name={key.description} on:delete={(event) => {
                                    modalOpen = true;
                                    deleteId = event.detail.keyId;
                               }}/>
                    {/each}
                </div>
            </Card>
        </div>
    </div>
    <ConfirmationModal description="This step will permanently delete the selected API key." on:cancel={() => {
       modalOpen = false;
       deleteId = "";
    }} on:proceed={async () => await deleteKey()} open={modalOpen}/>
</main>

{#if error !== undefined}
    <aside class="absolute left-10 {error ? 'come-up-animation' : 'come-down-animation'}">
        <Toast mode="error" size="md">
            {error}
        </Toast>
    </aside>
{/if}


<style lang="postcss">
    :root {
        --animation-delay: 0.25s;
    }

    @keyframes come-up {
        from {
            @apply -bottom-16;
        }
        to {
            @apply bottom-10 block;
        }
    }

    @keyframes come-down {
        from {
            @apply bottom-10;
        }
        to {
            @apply -bottom-16 hidden;
        }
    }

    .come-up-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-up;
        animation-fill-mode: forwards;
    }

    .come-down-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-down;
        animation-fill-mode: forwards;
    }
</style>

