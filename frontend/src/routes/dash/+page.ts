import type { PageLoad } from './$types';

export const load = (async () => {
    let respond = {}
    await fetch("/api/resources").then((res) => {
        if (res.ok) {
            respond = res.json();
        }
    })
    return respond;
}) satisfies PageLoad;