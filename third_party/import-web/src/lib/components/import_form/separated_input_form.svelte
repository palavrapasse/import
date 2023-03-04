<script lang="ts">
	export let id: string;
	export let label: string;
	export let name: string;
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

<label for={id}>{label}</label>
<input class="input input-primary input-sm" type="text" {name} {id} bind:value={innerValue} />

{#if value.length > 1}
	<div class="flex-row">
		{#each value as parcel}
			{#if parcel.length > 0}
				<span class="badge badge-primary">{parcel}</span>
			{/if}
		{/each}
	</div>
{/if}
