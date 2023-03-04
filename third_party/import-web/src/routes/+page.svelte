<script lang="ts">
	import { ImportForm, type Form } from '@components';

	function onSubmit(form: Form) {
		console.log(form);
	}
</script>

<h1>Import leak</h1>

<ImportForm on:submit={(event) => onSubmit(event.detail)} />
