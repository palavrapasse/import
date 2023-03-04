<script lang="ts">
	export let value: File;

	let files: FileList;

	$: {
		if (files?.length > 0) {
			value = files.item(0)!;
		}
	}
</script>

<label for="import-form-leak-file">Leak file</label>
<input type="file" name="leak file" id="import-form-leak file" bind:files />
