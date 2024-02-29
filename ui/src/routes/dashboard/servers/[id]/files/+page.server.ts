import type { PageServerLoad } from "./$types";
import { api } from "$lib/services/api";
import { getSignedInUser } from "$lib/services/auth_service";
import Client from "ssh2-sftp-client";

export const load: PageServerLoad = async ({ params }) => {
    params.id;
    const r = await api.serversService.findOneServers(params.id)
    const srv = r.response?.servers[0];
    const user = await getSignedInUser()

    const sftp = new Client();


    const SKYHOOK_HOST = process.env.SKYHOOK_HOST;
    await sftp.connect({
        host: SKYHOOK_HOST,
        port: 2022,
        username: user.id?.value,
        password: srv?.sftpPassword,
    }).catch(err => console.log(err));


    const path = process.env.SKYHOOK_STORAGE_LOCATION;
    const list = await sftp.list(`${path}`);

    console.log(list);

    return {
        files: list,
    }
}

