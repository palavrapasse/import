<script lang="ts">
	import SeparatedInputForm from './separated_input_form.svelte';

	export let value = new Array<string>();
</script>

<SeparatedInputForm
	id="import-form-platforms"
	label="Platforms (use ',' to include more than one platform)"
	name="leak platforms"
	bind:value
/>
