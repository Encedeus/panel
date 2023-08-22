<script>
    import Card from "$lib/components/generic/Card.svelte";
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import ProfilePictureIcon from "$lib/components/heroicons/ProfilePictureIcon.svelte";
    import ProfilePicture from "$lib/components/generic/ProfilePicture.svelte";
    import AccountIcon from "$lib/components/heroicons/AccountIcon.svelte";
    import AccountDetailTab from "$lib/components/internal/account/AccountDetailTab.svelte";
    import IdIcon from "$lib/components/heroicons/IdIcon.svelte";
    import MailIcon from "$lib/components/heroicons/MailIcon.svelte";
    import AccountChangeDetailModal from "$lib/components/internal/account/AccountChangeDetailModal.svelte";

    let changePasswordModalOpen = false;
    let changeUsernameModalOpen = false;
    let changeEmailModalOpen = false;
</script>

<main class="p-10 flex flex-col items-center">
    <div class="w-full h-full flex min-[1860px]:flex-row flex-col gap-8 items-center min-[1860px]:items-end">
        <div class="flex-grow">
            <CardHeader className="mb-5 self-start" size="lg">
                Account Settings
            </CardHeader>
            <Card className="w-[inherit] h-full" height="md" width="lg">
                <span class="text-sm" slot="title">
                    Account
                </span>
                <AccountIcon height={26} slot="icon" width={26}/>
                <div class="w-full h-full flex flex-col gap-3 items-center justify-between py-5 px-6" slot="content">
                    <AccountDetailTab on:change={() => changeUsernameModalOpen = true} label="Username" value="OptimusePrime">
                        <IdIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab on:change={() => changeEmailModalOpen = true} label="E-Mail" value="optimuseprime@optimuseprime.com">
                        <MailIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab on:change={() => changePasswordModalOpen = true} label="Password" value={new Array("1245678910".length).fill("*").join("")}>
                        <KeyIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                </div>
            </Card>
        </div>
        <div>
            <Card className="self-stretch" height="md" width="sm">
                <span class="text-sm" slot="title">
                    Profile Picture
                </span>
                <ProfilePictureIcon height={32} slot="icon" width={32}/>
                <div class="flex flex-col items-center justify-center h-full gap-4 p-8" slot="content">
                    <ProfilePicture changeable={true} height={375} width={375}/>
                </div>
            </Card>
        </div>
    </div>
    <AccountChangeDetailModal on:cancel={() => changePasswordModalOpen = false} subject="Password" open={changePasswordModalOpen}/>
    <AccountChangeDetailModal on:cancel={() => changeUsernameModalOpen = false} subject="Username" open={changeUsernameModalOpen}/>
    <AccountChangeDetailModal on:cancel={() => changeEmailModalOpen = false} subject="E-Mail" open={changeEmailModalOpen}/>
</main>

