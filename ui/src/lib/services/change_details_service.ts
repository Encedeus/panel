import type { HttpError, User } from "@encedeus/js-api";
import { api } from "$lib/services/api";
import {
    isWrongPasswordError,
    UserChangeEmailRequest,
    UserChangePasswordRequest,
    UserChangeUsernameRequest
} from "@encedeus/js-api";

export enum UserInformation {
    EMAIL = "email",
    PASSWORD = "password",
    USERNAME = "name",
}

export type AccountChangeDetails = {
    subject: UserInformation;
    oldSubject: string;
    newSubject: string;
    confirmNewSubject: string;
}

export type AccountChangeDetailResponse = {
    error: string;
    isInvalid: boolean;
};

export function subjectAsUppercase(subject: UserInformation): string | undefined {
    return subject?.toString().at(0)?.toUpperCase().concat(subject?.slice(1).toLowerCase());
}

export abstract class AccountChangeDetailService {

    constructor(protected user: User, protected details: AccountChangeDetails) {}

    protected isFieldEmpty(): AccountChangeDetailResponse | null {
        const isInvalid = !this.details.oldSubject.trim() || !this.details.newSubject.trim() || !this.details.confirmNewSubject.trim();
        if (isInvalid) {
            return {
                isInvalid: true,
                error: "Some fields are empty",
            };
        }

        return null;
    }

    protected isOldSubjectWrong(error: HttpError): AccountChangeDetailResponse | null {
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        const isInvalid = this.user[this.details.subject.toString()] !== this.details.oldSubject;
        if (!isInvalid) {
            return null;
        }

        return {
            isInvalid,
            error: `Old ${subjectAsUppercase(this.details.subject)} is wrong`,
        };
    }

    protected isNewSubjectEqualToOldSubject(): AccountChangeDetailResponse | null {
        const isInvalid = this.details.oldSubject === this.details.newSubject;
        if (!isInvalid) {
            return null;
        }

        return {
            isInvalid,
            error: `New ${subjectAsUppercase(this.details.subject)} must be different`,
        };
    }

    protected isNewSubjectNotEqualToConfirmation(): AccountChangeDetailResponse | null {
        const isInvalid = this.details.confirmNewSubject !== this.details.newSubject;
        if (!isInvalid) {
            return null;
        }

        return {
            isInvalid,
            error: `Conflict in new ${subjectAsUppercase(this.details.subject)}`,
        };
    }

    protected isAnyError(error: HttpError): AccountChangeDetailResponse | null {
        if (error) {
            return {
                error: error.message,
                isInvalid: true,
            };
        }

        return null;
    }

    protected validatePrefetchDetails(): AccountChangeDetailResponse | null {
        return this.isFieldEmpty() || this.isNewSubjectEqualToOldSubject() || this.isNewSubjectNotEqualToConfirmation();
    }

    protected validatePostfetchDetails(error: HttpError): AccountChangeDetailResponse | null {
        return this.isOldSubjectWrong(error) || this.isAnyError(error);
    }

    protected async sendChangeRequest(): Promise<HttpError | null> {
        return null;
    }

    async changeDetail(): Promise<AccountChangeDetailResponse | null> {
        const prefetchValidateResp = this.validatePrefetchDetails();
        if (prefetchValidateResp?.isInvalid) {
            return prefetchValidateResp;
        }

        const err = await this.sendChangeRequest();

        if (err) {
            const postfetchValidateResp = this.validatePostfetchDetails(err);
            if (postfetchValidateResp?.isInvalid) {
                return postfetchValidateResp;
            }
        }

        return null;
    }
}

export class AccountChangeEmailService extends AccountChangeDetailService {
    protected override async sendChangeRequest(): Promise<HttpError | null> {
        const resp = await api.usersService.changeEmail(UserChangeEmailRequest.create({
            oldEmail: this.details.oldSubject,
            newEmail: this.details.newSubject,
            userId: this.user.id,
        }));

        if (resp.error) {
            return resp.error;
        }

        return null;
    }
}

export class AccountChangeUsernameService extends AccountChangeDetailService {
    protected override async sendChangeRequest(): Promise<HttpError | null> {
        const resp = await api.usersService.changeUsername(UserChangeUsernameRequest.create({
            oldUsername: this.details.oldSubject,
            newUsername: this.details.newSubject,
            userId: this.user.id,
        }));

        if (resp.error) {
            return resp.error;
        }

        return null;
    }
}

export class AccountChangePasswordService extends AccountChangeDetailService {
    protected override async sendChangeRequest(): Promise<HttpError | null> {
        const resp = await api.usersService.changePassword(UserChangePasswordRequest.create({
            oldPassword: this.details.oldSubject,
            newPassword: this.details.newSubject,
            userId: this.user.id,
        }));

        if (resp.error) {
            return resp.error;
        }

        return null;
    }

    protected isOldSubjectWrong(error: HttpError): AccountChangeDetailResponse | null {
        if (isWrongPasswordError(error)) {
            return {
                isInvalid: true,
                error: `Old ${subjectAsUppercase(this.details.subject)} is wrong`,
            };
        }

        return null;
    }
}