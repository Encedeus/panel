import type { PageServerLoad } from "./$types";
import { api } from "$lib/services/api";
import { AccountAPIKeyFindManyByUserRequest, UUID } from "@encedeus/js-api";

export const load: PageServerLoad = async ({ locals }) => {
    const resp = await api.apiKeyService.findAccountApiKeysByUserId(AccountAPIKeyFindManyByUserRequest.create({
        userId: UUID.create({
            value: (locals as any).userId,
        }),
    }));

    return {
        apiKeys: resp.response?.accountApiKeys,
    };
};