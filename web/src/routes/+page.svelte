<script lang="ts">
    import { getWasm } from '$lib/wasm.js';
    import { browser } from '$app/environment';

    let message = $state('');
    let key = $state('');
    let output = $state('');
    let wasmLoaded = $state(false);
    let wasmLoading = $state(false);
    let wasmError = $state('');
    let gowasm = $state<Go | null>(null);

    // track load state
    if (browser) {
        wasmLoading = true;
        gowasm = getWasm()
            .then(() => { wasmLoaded = true; })
            .catch((e) => { wasmError = e.message; })
            .finally(() => { wasmLoading = false; });
    }

    function encrypt() {
        if (!wasmLoaded) return;
        if (!wasmLoaded) return;
        output = (window as any).wasmEncodeImage(message, key);
    }

    function decrypt() {
        if (!wasmLoaded) return;
        output = `Decrypted output will appear here`;
    }
</script>

<div class="page-wrapper">
    <div class="container">
        <!-- Header -->
        <div class="header">
            <h1>encryptr</h1>
            <p>Encrypt and decrypt messages with a key</p>
        </div>

        <!-- WASM status badge -->
        <div class="wasm-status">
            {#if wasmLoading}
                <span class="badge badge-loading">
                    <span class="spinner"></span>
                    Loading WASM…
                </span>
            {:else if wasmError}
                <span class="badge badge-error">
                    ✕ WASM failed — {wasmError}
                </span>
            {:else if wasmLoaded}
                <span class="badge badge-ready">
                    ● WASM ready
                </span>
            {/if}
        </div>

        <!-- Message field -->
        <div class="field">
            <label for="message">Message</label>
            <textarea
                id="message"
                bind:value={message}
                placeholder="Enter your message…"
                rows={4}
            ></textarea>
        </div>

        <!-- Key field -->
        <div class="field">
            <label for="key">Key</label>
            <input
                id="key"
                type="text"
                bind:value={key}
                placeholder="Enter encryption key…"
            />
        </div>

        <!-- Buttons -->
        <div class="buttons">
            <button
                class="btn btn-primary"
                onclick={encrypt}
                disabled={!message || !key || !wasmLoaded}
            >
                Encrypt
            </button>
            <button
                class="btn btn-secondary"
                onclick={decrypt}
                disabled={!message || !key || !wasmLoaded}
            >
                Decrypt
            </button>
        </div>

        <!-- Output -->
        {#if output}
            <div class="output">
                <span class="output-label">Output</span>
                <div class="output-box">{output}</div>
            </div>
        {/if}
    </div>
</div>