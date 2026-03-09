// src/lib/wasm.js
import { browser } from '$app/environment';

let wasmPromise: Promise<void> | null = null;

export async function getWasm() {
    if (!browser) return null;
    
    if (!wasmPromise) {
        wasmPromise = await initWasm();
    }
    
    return wasmPromise;
}

async function initWasm() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
        fetch('/main.wasm'),
        go.importObject
    );
    go.run(result.instance);
}