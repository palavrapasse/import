<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Context from './context.svelte';
	import Leakers from './leakers.svelte';
	import LeakFile from './leak_file.svelte';
	import Platforms from './platforms.svelte';
	import ShareDate from './share_date.svelte';
	import SubmitButton from './submit_button.svelte';
	import type { Form } from './types';

	const dispatcher = createEventDispatcher();
	let form = <Form>{};

	function submitForm() {
		dispatcher('submit', form);
	}
</script>

<form class="flex flex-col gap-2" on:submit|preventDefault={submitForm} action=".">
	<Context bind:value={form.context} />
	<ShareDate bind:value={form.shareDateMS} />
	<Platforms bind:value={form.platforms} />
	<Leakers bind:value={form.leakers} />
	<LeakFile bind:value={form.leakFile} />

	<SubmitButton />
</form>
