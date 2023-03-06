<script lang="ts">
	import { ImportForm, type Form } from '@components';
	import http from '@http';

	let response: Response | Error;

	function onSubmit(form: Form) {
		const request = http(form);

		request.then((resp) => (response = resp)).catch((err) => (response = new Error(err)));
	}
</script>

<div class="flex flex-row justify-center">
	<div class="card max-w-fit">
		<div class="card-body">
			<h2 class="card-header">Import leak</h2>
			<ImportForm on:submit={(event) => onSubmit(event.detail)} />

			{#if response instanceof Response}
				<p class="text-success text-center">Uploaded!</p>
			{:else if response instanceof Error}
				<p class="text-error text-center">Failed to upload! ({response})</p>
			{/if}
		</div>
	</div>
</div>
