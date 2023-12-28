import type { LayoutServerLoad } from "./$types";
import axios from "axios";

export const load: LayoutServerLoad = async () => {
    const resp = await axios.get("http://localhost:8080/modules");

    return {
        modules: resp.data.modules as any[],
    };
};