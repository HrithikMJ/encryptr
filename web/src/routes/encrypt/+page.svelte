<script lang="ts">
	import { getContext } from 'svelte';

	const wasm: { loaded: boolean; loading: boolean; error: string } = getContext('wasm');

	let message = $state('');
	let key = $state('');
	let output = $state('');
	let encryptError = $state('');

	async function encrypt() {
		if (!wasm.loaded) return;
		encryptError = '';
		try {
			const base64 = await (window as any).wasmEncodeImage(message, key);
			output = 'data:image/png;base64,' + base64;
		} catch (err: any) {
			output = '';
			encryptError = typeof err === 'string' ? err : err.message ?? String(err);
		}
	}

	function download() {
		if (!output) return;
		const a = document.createElement('a');
		a.href = output;
		a.download = 'encrypted.png';
		a.click();
	}
</script>

<div class="page-wrapper">
	<div class="container">
		<div class="header">
			<div class="header-row">
				<a href="/" class="back-link">← Back</a>
			</div>
			<h1>Encrypt</h1>
			<p>Enter a message and key to generate an encrypted PNG</p>
		</div>

		{#if encryptError}
			<div class="error-banner">{encryptError}</div>
		{/if}

		<!-- Message field -->
		<div class="field">
			<label for="message">Message</label>
			<textarea id="message" bind:value={message} placeholder="Enter your message…" rows={4}
			></textarea>
		</div>

		<!-- Key field -->
		<div class="field">
			<label for="key">Key</label>
			<input id="key" type="text" bind:value={key} placeholder="Enter encryption key…" />
		</div>

		<!-- Button -->
		<div class="buttons">
			<button class="btn btn-primary" onclick={encrypt} disabled={!message || !key || !wasm.loaded}>
				Encrypt
			</button>
		</div>

		<!-- Output -->
		{#if output}
			<div class="output">
				<span class="output-label">Output</span>
				<div class="output-preview">
					<img src={output} alt="Encrypted output" />
				</div>
				<button class="btn btn-primary" onclick={download}> Download PNG </button>
			</div>
		{/if}
	</div>
</div>

<style>
	.header-row {
		margin-bottom: 0.75rem;
	}
	.back-link {
		font-size: 0.8125rem;
		color: #71717a;
		text-decoration: none;
		transition: color 0.15s ease;
	}
	.back-link:hover {
		color: #a1a1aa;
	}
	.error-banner {
		padding: 0.625rem 1rem;
		font-size: 0.8125rem;
		color: #f87171;
		background-color: #450a0a;
		border: 1px solid #7f1d1d;
		border-radius: 0.5rem;
		margin-bottom: 1.25rem;
	}
</style>
