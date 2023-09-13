<script lang="ts">
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
    import { onMount } from "svelte";
    import { getSignedInUser } from "$lib/services/auth_service.js";
    import { User } from "@encedeus/js-api";
    import Toast from "$lib/components/generic/Toast.svelte";
    import type { AccountChangeDetails } from "$lib/services/change_details_service";
    import {
        AccountChangeDetailService,
        AccountChangeEmailService,
        AccountChangePasswordService,
        AccountChangeUsernameService,
        subjectAsUppercase,
        UserInformation
    } from "$lib/services/change_details_service";

    let user = User.create();

    let notification: string | null | undefined = undefined;
    let notificationMode: "ok" | "error" = "ok";

    let changeModalOpen = false;
    let changeDetails: AccountChangeDetails = {};
    let oldSubjectError = false;
    let newSubjectError = false;
    let confirmNewSubjectError = false;

    onMount(async () => {
        user = await getSignedInUser();
    });


    async function onChangeDetail() {
        let service: AccountChangeDetailService;

        switch (changeDetails.subject) {
        case UserInformation.EMAIL:
            service = new AccountChangeEmailService(user, changeDetails);
            break;
        case UserInformation.PASSWORD:
            service = new AccountChangePasswordService(user, changeDetails);
            break;
        case UserInformation.USERNAME:
            service = new AccountChangeUsernameService(user, changeDetails);
            break;
        }

        const resp = await service.changeDetail();
        if (resp?.isInvalid) {
            errorNotification(resp?.error);
            return;
        }

        user[changeDetails?.subject] = changeDetails.newSubject;

        changeModalOpen = false;
        okNotification(`Changed ${subjectAsUppercase(changeDetails?.subject)} successfully`);
        changeDetails = {};
    }

    function errorNotification(notify: string, timeout?: boolean) {
        notification = notify;
        notificationMode = "error";
        if (timeout) {
            setTimeout(() => {
                notification = "";
            }, 2000);
        }
    }

    function okNotification(notify: string) {
        notification = notify;
        notificationMode = "ok";
        setTimeout(() => {
            notification = "";
        }, 2000);
    }

    function clearNotification() {
        oldSubjectError = false;
        newSubjectError = false;
        confirmNewSubjectError = false;

        if (notification) {
            notification = "";
            notificationMode = "error";
        }
    }
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
                    <AccountDetailTab bind:value={user.name} label="Username" on:change={() => {
                        changeModalOpen = true;
                        changeDetails.subject = UserInformation.USERNAME;
                        changeDetails = changeDetails;
                    }}>
                        <IdIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab bind:value={user.email} label="E-Mail" on:change={() => {
                        changeModalOpen = true;
                        changeDetails.subject = UserInformation.EMAIL;
                        changeDetails = changeDetails;
                    }}>
                        <MailIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab label="Password" on:change={() => {
                        changeModalOpen = true;
                        changeDetails.subject = UserInformation.PASSWORD;
                        changeDetails = changeDetails;
                    }} value={new Array(12).fill("*").join("")}>
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
    {#if notification !== undefined}
        <aside class="absolute left-10 {notification ? 'come-up-animation' : 'come-down-animation'}">
            <Toast mode={notificationMode} size="md">
                {notification}
            </Toast>
        </aside>
    {/if}

    <AccountChangeDetailModal bind:confirmNewSubjectError={confirmNewSubjectError} bind:newSubjectError={newSubjectError}
                              bind:oldSubjectError={oldSubjectError}
                              on:cancel={() => { changeModalOpen = false; clearNotification(); }}
                              on:input={clearNotification} on:save={onChangeDetail} open={changeModalOpen}
                              subjectDetails={changeDetails}/>
</main>

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
