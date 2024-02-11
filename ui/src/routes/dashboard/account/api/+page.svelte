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
    import { isIP } from "is-ip";
    import { onMount } from "svelte";
    import CursorArrowRays from "$lib/components/heroicons/CursorArrowRays.svelte";
    import {
        AccountAPIKey,
        AccountAPIKeyDeleteRequest,
        UUID,
        AccountAPIKeyCreateRequest
    } from "@encedeus/js-api";
    import type { PageServerData } from "./$types";
    import { fly } from "svelte/transition";
    import Notification from "$lib/components/generic/Notification.svelte";

    let modalOpen = false;

    let keyDescription = "";
    let allowedIpsBox = "";

    let notification: string | undefined | null = undefined;
    let notificationMode: "error" | "ok" = "error";
    let descriptionError = false;
    let ipError = false;

    let candidateForDeletion = "";

    let apiKeys: AccountAPIKey[] = [];
    export let data: PageServerData;
    $: apiKeys = data.apiKeys ?? [];

    onMount(async () => {
        if (!apiKeys && apiKeys?.length === 0) {
            errorNotification("Failed fetching API keys");
        }
    });

    function validateInput(): boolean {
        if (!keyDescription) {
            errorNotification("No description provided");
            descriptionError = true;
            return false;
        }
        if (keyDescription.length > 28) {
            errorNotification("Description is too long");
            descriptionError = true;
            return false;
        }
        if (keyDescription.length < 3) {
            errorNotification("Description is too short");
            descriptionError = true;
            return false;
        }

        for (const ip of allowedIpsBox.trim().split("\n")) {
            if (!isIP(ip) && ip) {
                errorNotification("Invalid IP address");
                ipError = true;
                return false;
            }
        }

        return true;
    }

    async function createKey() {
        if (notification) {
            notification = "";
        }

        if (!validateInput()) {
            return;
        }

        const key = AccountAPIKey.create();
        key.description = keyDescription;
        key.ipAddresses = allowedIpsBox.trim().split("\n");
        key.userId = UUID.create((await getSignedInUser()).id);

        const { error, response } = await api.apiKeyService.createAccountApiKey(AccountAPIKeyCreateRequest.create({
            ipAddresses: key.ipAddresses,
            description: key.description,
            userId: key.userId,
        }));
        if (error) {
            notification = error.message;
            return;
        }
        if (!response) {
            errorNotification("Something went wrong");
            return;
        }

        key.key = <string>response?.accountApiKey?.key;
        key.id = response?.accountApiKey?.id;
        apiKeys = [key, ...apiKeys];

        okNotification("Created an API key");

        keyDescription = "";
        allowedIpsBox = "";
    }

    async function deleteKey(keyId: UUID) {
        const idx = apiKeys.findIndex(key => key.id?.value === keyId.value);
        apiKeys.splice(idx, 1);
        apiKeys = apiKeys;

        const resp = await api.apiKeyService.deleteAccountApiKey(AccountAPIKeyDeleteRequest.create({
            id: UUID.create(keyId),
        }));
        if (resp.error) {
            errorNotification("Failed deleting API key");
            return;
        }

        okNotification("Deleted an API key");
    }

    async function onDelete() {
        modalOpen = false;
        if (!candidateForDeletion) {
            errorNotification("Something went wrong");
            return;
        }
        await deleteKey(candidateForDeletion);
    }

    function clearInput() {
        if (notification) {
            notification = "";
            descriptionError = false;
            ipError = false;
        }
    }

    function errorNotification(notify: string, timeout?: boolean) {
        notification = notify;
        notificationMode = "error";
        if (timeout) {
            setTimeout(() => {
                notification = "";
            }, 1500);
        }
    }

    function okNotification(notify: string) {
        notification = notify;
        notificationMode = "ok";
        setTimeout(() => {
            notification = "";
        }, 1500);
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
<!--                <div class="mt-5 w-full h-full flex flex-col items-center justify-between" slot="content">-->
                    <div class="mt-7 w-full h-full flex flex-col items-center basis-0 flex-grow justify-center gap-4" slot="content">
                        <Input bind:value={keyDescription} error={descriptionError} label="Description"
                               on:input={clearInput}
                               placeholder="API Key Description" size="xl"/>
                        <TextArea bind:value={allowedIpsBox} className="basis-full resize-none" error={ipError}
                                  label="Allowed IPs"
                                  on:input={clearInput}
                                  placeholder="Leave blank to allow any IP address to use this API key. Otherwise, provide each IP address on a new line."
                                  size="lg"/>
                        <Button className="mt-7 flex justify-center items-center" color="indigo"
                                on:click={async () => await createKey()}
                                size="sm">Create
                        </Button>
                    </div>
<!--                </div>-->
            </Card>
        </div>
        <div>
            <Card fixedHeight={true} height="md" width="lg">
                <span class="text-sm" slot="title">
                    API Keys
                </span>
                <KeyIcon slot="icon"/>
                <div class="flex flex-col items-center justify-center gap-4 p-8" slot="content">
                    {#if apiKeys && apiKeys.length > 0}
                        {#each apiKeys as key}
                            <KeyTab className="flex-grow w-full h-full" id={key.id} key={key.key}
                                    name={key.description}
                                    on:delete={() => { modalOpen = true; candidateForDeletion = key.id; console.log(key.id); }}/>
                        {/each}
                    {:else}
                        <CursorArrowRays width={240} height={240}/>
                        <span class="text-indigo-600 text-lg">Nothing here yet. Create a new API key to see something!</span>
                    {/if}
                </div>
            </Card>
        </div>
    </div>
    <ConfirmationModal bind:open={modalOpen}
                       description="This step will permanently delete the selected API key."
                       on:cancel={() => { modalOpen = false; candidateForDeletion = ""; }}
                       on:proceed={async () => await onDelete()}/>
</main>

<Notification show={notification} {notification} mode={notificationMode}/>



