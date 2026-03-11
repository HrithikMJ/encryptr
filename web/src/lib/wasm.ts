// src/lib/wasm.js
import { browser } from '$app/environment';
import { getContext } from 'svelte';

let wasmPromise: Promise<void> | null = null;

export async function getWasm() {
    if (!browser) return null;
    
    if (!wasmPromise) {
        wasmPromise = initWasm();
    }
    
    return wasmPromise;
}

async function initWasm() {
    const go = new Go();
    const wasm: { loaded: boolean; loading: boolean; error: string } = getContext('wasm');
    const result = await WebAssembly.instantiateStreaming(
        fetch('/main.wasm'),
        go.importObject
    );
    go.run(result.instance).then((err: any) => {
        console.log('WASM initialized');
        wasm.error = 'Go program exited';
    }).catch((error) => {
        console.error('WASM error:', error);
        wasm.error = error.message ?? String(error);
    });
}