import type { PageServerLoad } from "./$types";
import { api } from "$lib/services/api";
import { getSignedInUser } from "$lib/services/auth_service";
import Client from "ssh2-sftp-client";

export const load: PageServerLoad = async ({ url, params, locals }) => {
    const r = await api.serversService.findOneServers(params.id)
    const srv = r.response?.servers[0];

    const sftp = new Client();

    await sftp.connect({
        host: "127.0.0.1",
        port: 2022,
        username: (locals as any).userId,
        password: srv?.sftpPassword,
    }).catch(err => console.log(err));

    const path = `${process.env.SKYHOOK_STORAGE_LOCATION}/${params.path}`;
    const list = await sftp.list(`${path}`);

    return {
        files: list,
    }
}

