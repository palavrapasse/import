<script lang="ts">
	import { ImportForm, type Form, LoadingSpinner } from '@components';
	import http from '@http';

	let isProcessingRequest = false;
	let response: Response | Error;
	let responseBody: String;
	let leaksDbSizeChangeInBytes: number | undefined;

	function onSubmit(form: Form) {
		isProcessingRequest = true;

		const request = http(form);

		request
			.then((resp) => (response = resp))
			.then((resp) => resp.text())
			.then((respBody) => {
				responseBody = respBody;
				leaksDbSizeChangeInBytes = Number.parseFloat(respBody);
			})
			.catch((err) => (response = new Error(err)))
			.finally(() => (isProcessingRequest = false));
	}
</script>

<div class="flex flex-row justify-center">
	<div class="card max-w-fit">
		<div class="card-body">
			<h2 class="card-header">Import leak</h2>
			<ImportForm on:submit={(event) => onSubmit(event.detail)} />

			{#if isProcessingRequest}
				<div class="self-center">
					<LoadingSpinner />
				</div>
			{:else if response instanceof Response}
				{#if leaksDbSizeChangeInBytes !== undefined && !isNaN(leaksDbSizeChangeInBytes)}
					{#if leaksDbSizeChangeInBytes > 0}
						<p class="text-success text-center">Uploaded!</p>
						<p class="text-center">
							LeaksDB size increase: {leaksDbSizeChangeInBytes / (1024 * 1024)} MB
						</p>
					{:else}
						<p class="text-error text-center">Failed to import: LeaksDB size didn't change.</p>
					{/if}
				{:else}
					<p class="text-error text-center">
						Failed to import: {responseBody}
					</p>
				{/if}
				<p class="text-center" />
			{:else if response instanceof Error}
				<p class="text-error text-center">Failed to upload! ({response})</p>
			{/if}
		</div>
	</div>
</div>
