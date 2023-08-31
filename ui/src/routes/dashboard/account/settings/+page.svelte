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
    import {
        User,
        UserChangeUsernameRequest,
        UserChangeEmailRequest,
        UserChangePasswordRequest,
        isWrongPasswordError
    } from "@encedeus/js-api";
    import { api } from "$lib/services/api";
    import Toast from "$lib/components/generic/Toast.svelte";

    let changePasswordModalOpen = false;
    let changeUsernameModalOpen = false;
    let changeEmailModalOpen = false;

    let user: User;

    let oldUsername = "";
    let oldUsernameError = false;
    let newUsername = "";
    let newUsernameError = false;
    let confirmNewUsername = "";
    let confirmNewUsernameError = false;

    let oldEmail = "";
    let oldEmailError = false;
    let newEmail = "";
    let newEmailError = false;
    let confirmNewEmail = "";
    let confirmNewEmailError = false;

    let oldPassword = "";
    let oldPasswordError = false;
    let newPassword = "";
    let newPasswordError = false;
    let confirmNewPassword = "";
    let confirmNewPasswordError = false;

    let notification: string | null | undefined = undefined;
    let notificationMode: "ok" | "error" = "ok";

    onMount(async () => {
        user = await getSignedInUser();
    });

    async function onChangeUsername() {
        if (!oldUsername.trim() || !newUsername.trim() || !confirmNewUsername.trim()) {
            errorNotification("You must fill all fields");
            return;
        }
        if (oldUsername.trim() !== user.name) {
            oldUsernameError = true;
            errorNotification("Old username is wrong");
            return;
        }
        if (newUsername.trim() === oldUsername.trim()) {
            newUsernameError = true;
            errorNotification("New username must be different");
            return;
        }
        if (newUsername !== confirmNewUsername) {
            newUsernameError = true;
            confirmNewUsernameError = true;
            errorNotification("Conflict in new username");
            return;
        }

        const { error } = await api.usersService.changeUsername(UserChangeUsernameRequest.create({
            userId: user.id,
            oldUsername,
            newUsername,
        }));
        if (error) {
            errorNotification("Something went wrong");
            return;
        }

        user.name = newUsername;
        changeUsernameModalOpen = false;
        oldUsername = "";
        newUsername = "";
        confirmNewUsername = "";

        okNotification("Changed username successfully");
    }

    async function onChangeEmail() {
        if (!oldEmail.trim() || !newEmail.trim() || !confirmNewEmail.trim()) {
            errorNotification("You must fill all fields");
            return;
        }
        if (oldEmail.trim() !== user.email) {
            oldEmailError = true;
            errorNotification("Old email is wrong");
            return;
        }
        if (newEmail.trim() === oldEmail.trim()) {
            newEmailError = true;
            errorNotification("New email must be different");
            return;
        }
        if (newEmail !== confirmNewEmail) {
            newEmailError = true;
            confirmNewEmailError = true;
            errorNotification("Conflict in new email");
            return;
        }

        const { error } = await api.usersService.changeEmail(UserChangeEmailRequest.create({
            userId: user.id,
            oldEmail,
            newEmail,
        }));
        if (error) {
            errorNotification("Something went wrong");
            return;
        }

        user.email = newEmail;
        changeEmailModalOpen = false;
        oldEmail = "";
        newEmail = "";
        confirmNewEmail = "";

        okNotification("Changed email successfully");
    }

    async function onChangePassword() {
        if (!oldPassword.trim() || !newPassword.trim() || !confirmNewPassword.trim()) {
            errorNotification("You must fill all fields");
            return;
        }
        if (newPassword !== confirmNewPassword) {
            newPasswordError = true;
            confirmNewPasswordError = true;
            errorNotification("Conflict in new password");
            return;
        }

        const { error } = await api.usersService.changePassword(UserChangePasswordRequest.create({
            userId: user.id,
            oldPassword,
            newPassword,
        }));
        if (error) {
            if (isWrongPasswordError(error)) {
                errorNotification("Old password is wrong");
                return;
            }
            errorNotification("Something went wrong");
            return;
        }

        changePasswordModalOpen = false;
        oldPassword = "";
        newPassword = "";
        confirmNewPassword = "";

        okNotification("Changed password successfully");
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
        oldUsernameError = false;
        newUsernameError = false;
        confirmNewUsernameError = false;
        oldEmailError = false;
        newEmailError = false;
        confirmNewEmailError = false;
        oldPasswordError = false;
        newPasswordError = false;
        confirmNewPasswordError = false;
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
                    <AccountDetailTab on:change={() => changeUsernameModalOpen = true} label="Username" value={user?.name}>
                        <IdIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab on:change={() => changeEmailModalOpen = true} label="E-Mail" value={user?.email}>
                        <MailIcon height={32} slot="icon" width={32}/>
                    </AccountDetailTab>
                    <AccountDetailTab on:change={() => changePasswordModalOpen = true} label="Password" value={new Array(12).fill("*").join("")}>
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

    <AccountChangeDetailModal bind:oldSubjectError={oldPasswordError} bind:newSubjectError={newPasswordError} bind:confirmNewSubjectError={confirmNewPasswordError} on:cancel={() => { changePasswordModalOpen = false; clearNotification(); }} on:input={clearNotification} on:save={onChangePassword} bind:oldSubject={oldPassword} bind:newSubject={newPassword} bind:confirmNewSubject={confirmNewPassword} subject="Password" open={changePasswordModalOpen}/>
    <AccountChangeDetailModal bind:oldSubjectError={oldUsernameError} bind:newSubjectError={newUsernameError} bind:confirmNewSubjectError={confirmNewUsernameError} on:cancel={() => { changeUsernameModalOpen = false; clearNotification(); }} on:input={clearNotification} on:save={onChangeUsername} bind:oldSubject={oldUsername} bind:newSubject={newUsername} bind:confirmNewSubject={confirmNewUsername} subject="Username" open={changeUsernameModalOpen}/>
    <AccountChangeDetailModal bind:oldSubjectError={oldEmailError} bind:newSubjectError={newEmailError} bind:confirmNewSubjectError={confirmNewEmailError} on:cancel={() => { changeEmailModalOpen = false; clearNotification(); }} on:input={clearNotification} on:save={onChangeEmail} bind:oldSubject={oldEmail} bind:newSubject={newEmail} bind:confirmNewSubject={confirmNewEmail} subject="E-Mail" open={changeEmailModalOpen}/>
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
