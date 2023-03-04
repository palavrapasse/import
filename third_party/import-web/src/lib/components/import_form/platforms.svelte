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

<label for="import-form-platforms">Platforms (use "," to include more than one platform)</label>
<input
	class="input input-primary input-sm"
	type="text"
	name="leak platforms"
	id="import-form-platforms"
	bind:value={innerValue}
/>
