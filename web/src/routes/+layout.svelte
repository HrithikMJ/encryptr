<script lang="ts">
	import '../app.css';
	import { getWasm } from '$lib/wasm.js';
	import { browser } from '$app/environment';
	import { setContext } from 'svelte';
	
	setContext('wasm', {
		get loaded() { return wasmLoaded; },
		get loading() { return wasmLoading; },
		get error() { return wasmError; },
		set error(error: string) { wasmError = error; }
	});
	let { children } = $props();
	
	let wasmLoaded = $state(false);
	let wasmLoading = $state(false);
	let wasmError = $state('');
	let alive = true;

	if (browser) {
		wasmLoading = true;
		getWasm()
			.then(() => {
				wasmLoaded = true;
			})
			.catch((e: any) => {
				wasmError = e.message ?? String(e);
			})
			.finally(() => {
				wasmLoading = false;
			});
	}

</script>

<!-- WASM status badge (shown on all pages) -->
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
		<span class="badge badge-ready"> ● WASM ready </span>
	{/if}
</div>

{@render children()}
