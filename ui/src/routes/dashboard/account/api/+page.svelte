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

    let modalOpen = false;
</script>

<main class="p-6 flex flex-col items-center">
    <div class="flex min-[1860px]:flex-row flex-col gap-8 items-end justify-end">
        <div class="w-full h-full">
            <CardHeader size="lg" className="mb-5 self-start">
                API Credentials
            </CardHeader>
            <Card width="lg" height="md" className="w-full h-full">
                <span slot="title" class="text-sm">
                    Create API Key
                </span>
                <PlusIcon slot="icon"/>
                <div slot="content" class="mt-5 w-full h-full flex flex-col items-center justify-between">
                    <div class="w-full h-full flex flex-col items-center justify-center gap-4">
                        <Input size="xl" label="Description" placeholder="API Key Description"/>
                        <TextArea size="lg" label="Allowed IPs" placeholder="Leave blank to allow any IP address to use this API key. Otherwise, provide each IP address on a new line." className="basis-full resize-none"/>
                        <Button size="sm" color="indigo" className="mt-6 flex justify-center items-center">Create</Button>
                    </div>
                </div>
            </Card>
        </div>        
        <div>
            <Card width="lg" height="md" fixedHeight={true}>
                <span slot="title" class="text-sm">
                    API Keys
                </span>
                <KeyIcon slot="icon"/>
                <div slot="content" class="flex flex-col items-center justify-center h-full gap-4 p-8">
                    <KeyTab on:delete={() => modalOpen = true} name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                    <KeyTab name="Neovim Key" lastUsed="AUG 25TH, 2020 19:20" key="NfFv11gxmcmP5faS" className="flex-grow w-full"/>
                </div>
            </Card>
        </div>
    </div>
    <ConfirmationModal open={modalOpen} on:cancel={() => modalOpen = false} description="This step will permanently delete the selected API key."/>
</main>

