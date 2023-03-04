<script lang="ts">
	export const value = new Array<string>();
	let innerValue = '';

	$: {
		const trimmed = innerValue.replaceAll(/\\s+|,*/g, '');
		let currentIndex = value.length - 1;

		if (currentIndex < 0) {
			currentIndex = 0;
		}

		if (trimmed.length > 1 && innerValue.endsWith(',')) {
			innerValue = '';
			currentIndex++;
		}

		value[currentIndex] = innerValue;
	}
</script>

<label for="import-form-leakers">Leakers (use "," to include more than one leaker)</label>
<input type="text" name="leak bad actors" id="import-form-leakers" bind:value={innerValue} />
