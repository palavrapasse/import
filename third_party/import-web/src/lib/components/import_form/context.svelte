<script lang="ts">
	export let value: string;
</script>

<label for="import-form-context">Context</label>
<input
	class="input input-primary input-sm"
	type="text"
	name="leak context"
	id="import-form-context"
	bind:value
/>
