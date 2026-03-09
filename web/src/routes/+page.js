import { browser } from '$app/environment';

export async function load() {
    if (browser) {
        const { getWasm } = await import('$lib/wasm.js');
        await getWasm();
    }
    return {};
}