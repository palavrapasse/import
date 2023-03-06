<script lang="ts">
	export let value: number;
	let innerValue: string = '1970-01-01';

	$: value = Date.parse(innerValue);
</script>

<label for="import-form-share-date">Share date (defaults to 01-01-1970/UNKNOWN)</label>
<input
	class="input input-primary input-sm"
	type="date"
	name="leak share date"
	id="import-form-share-date"
	bind:value={innerValue}
/>
