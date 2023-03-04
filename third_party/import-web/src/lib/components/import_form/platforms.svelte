<script lang="ts">
	export const value = new Array<string>();
	let innerValue = '';

	$: {
		const trimmed = innerValue.replaceAll(/\\s+|,*/g, '');

		if (trimmed.length > 1 && innerValue.endsWith(',')) {
			value.push(trimmed);
			innerValue = '';
		}
	}
</script>

<label for="import-form-platforms">Platforms (use "," to include more than one platform)</label>
<input type="text" name="leak platforms" id="import-form-platforms" bind:value={innerValue} />
