<script lang="ts">
	import { ImportForm, type Form } from '@components';

	function onSubmit(form: Form) {
		console.log(form);
	}
</script>

<div class="flex flex-row justify-center">
	<div class="card max-w-fit">
		<div class="card-body">
			<h2 class="card-header">Import leak</h2>
			<ImportForm on:submit={(event) => onSubmit(event.detail)} />
		</div>
	</div>
</div>
