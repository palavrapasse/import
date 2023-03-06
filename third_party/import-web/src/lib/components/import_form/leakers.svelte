<script lang="ts">
	import SeparatedInputForm from './separated_input_form.svelte';

	export let value = new Array<string>();
</script>

<SeparatedInputForm
	id="import-form-leakers"
	label="Leakers (use ',' to include more than one leaker)"
	name="leak bad actors"
	bind:value
/>
