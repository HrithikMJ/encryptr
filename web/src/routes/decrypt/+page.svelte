<script lang="ts">
	import { getContext } from 'svelte';

	const wasm: { loaded: boolean; loading: boolean; error: string } = getContext('wasm');

	let key = $state('');
	let output = $state('');
	let decryptError = $state('');
	let imagePreview = $state('');
	let imageBase64 = $state('');
	let dragOver = $state(false);
	let fileName = $state('');
	let fileInput = $state<HTMLInputElement>();

	function handleFile(file: File) {
		if (!file.type.startsWith('image/')) {
			decryptError = 'Please upload a valid image file (PNG)';
			return;
		}
		decryptError = '';
		fileName = file.name;
		const reader = new FileReader();
		reader.onload = () => {
			const dataUrl = reader.result as string;
			imagePreview = dataUrl;
			// Strip the data URL prefix to get pure base64
			imageBase64 = dataUrl.replace(/^data:image\/\w+;base64,/, '');
		};
		reader.readAsDataURL(file);
	}

	function onFileInput(e: Event) {
		const input = e.target as HTMLInputElement;
		if (input.files?.[0]) {
			handleFile(input.files[0]);
		}
	}

	function onDrop(e: DragEvent) {
		e.preventDefault();
		dragOver = false;
		if (e.dataTransfer?.files?.[0]) {
			handleFile(e.dataTransfer.files[0]);
		}
	}

	function onDragOver(e: DragEvent) {
		e.preventDefault();
		dragOver = true;
	}

	function onDragLeave() {
		dragOver = false;
	}

	function onDropZoneClick() {
		fileInput?.click();
	}

	function onDropZoneKeyDown(e: KeyboardEvent) {
		if (e.key === 'Enter' || e.key === ' ') {
			e.preventDefault();
			fileInput?.click();
		}
	}

	function clearImage() {
		imagePreview = '';
		imageBase64 = '';
		fileName = '';
		output = '';
		decryptError = '';
	}

	async function decrypt() {
		if (!wasm.loaded || !imageBase64 || !key) return;
		decryptError = '';
		output = '';
		try {
			const plaintext = await (window as any).wasmDecodeImage(imageBase64, key);
			output = plaintext;
		} catch (err: any) {
			decryptError = typeof err === 'string' ? err : err.message ?? String(err);
		}
	}
</script>

<div class="page-wrapper">
	<div class="container">
		<div class="header">
			<div class="header-row">
				<a href="/" class="back-link">← Back</a>
			</div>
			<h1>Decrypt</h1>
			<p>Upload an encrypted PNG and provide the key to reveal the message</p>
		</div>

		{#if decryptError}
			<div class="error-banner">{decryptError}</div>
		{/if}

		<!-- Image upload -->
		<div class="field">
			{#if imagePreview}
				<div class="image-uploaded">
					<div class="output-preview">
						<img src={imagePreview} alt="Uploaded encrypted" />
					</div>
					<div class="image-info">
						<span class="image-name">{fileName}</span>
						<button class="btn-remove" onclick={clearImage}>Remove</button>
					</div>
				</div>
			{:else}
			<div
					class="drop-zone"
					class:drop-zone-active={dragOver}
					role="button"
					tabindex="0"
					aria-label="Drop zone — drag and drop an image or press Enter to browse"
					ondrop={onDrop}
					ondragover={onDragOver}
					ondragleave={onDragLeave}
					onclick={onDropZoneClick}
					onkeydown={onDropZoneKeyDown}
				>
					<div class="drop-zone-content">
						<span class="drop-icon">📁</span>
						<span class="drop-text">Drag & drop an image here</span>
						<span class="drop-or">or</span>
						<label class="btn btn-secondary upload-btn">
							Browse files
							<input
								bind:this={fileInput}
								type="file"
								accept="image/png,image/*"
								onchange={onFileInput}
								hidden
							/>
						</label>
					</div>
				</div>
			{/if}
		</div>

		<!-- Key field -->
		<div class="field">
			<label for="key">Key</label>
			<input id="key" type="text" bind:value={key} placeholder="Enter decryption key…" />
		</div>

		<!-- Button -->
		<div class="buttons">
			<button
				class="btn btn-primary"
				onclick={decrypt}
				disabled={!imageBase64 || !key || !wasm.loaded}
			>
				Decrypt
			</button>
		</div>

		<!-- Output -->
		{#if output}
			<div class="output">
				<span class="output-label">Decrypted message</span>
				<div class="output-box">{output}</div>
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

	/* Drop zone */
	.drop-zone {
		border: 2px dashed #27272a;
		border-radius: 0.75rem;
		padding: 2.5rem 1rem;
		text-align: center;
		transition: border-color 0.15s ease, background-color 0.15s ease;
		cursor: pointer;
	}
	.drop-zone:hover,
	.drop-zone-active {
		border-color: #52525b;
		background-color: #18181b;
	}
	.drop-zone-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
	}
	.drop-icon {
		font-size: 2rem;
		opacity: 0.6;
	}
	.drop-text {
		font-size: 0.875rem;
		color: #a1a1aa;
	}
	.drop-or {
		font-size: 0.75rem;
		color: #52525b;
	}
	.upload-btn {
		cursor: pointer;
		flex: unset;
		padding: 0.5rem 1rem;
	}

	/* Uploaded image preview */
	.image-uploaded {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}
	.image-info {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.image-name {
		font-size: 0.8125rem;
		color: #71717a;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}
	.btn-remove {
		font-size: 0.75rem;
		color: #f87171;
		background: none;
		border: none;
		cursor: pointer;
		padding: 0.25rem 0.5rem;
		border-radius: 0.25rem;
		transition: background-color 0.15s ease;
	}
	.btn-remove:hover {
		background-color: #450a0a;
	}
</style>
