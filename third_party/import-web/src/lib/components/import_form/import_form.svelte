<script lang="ts">
	import Context from './context.svelte';
	import Leakers from './leakers.svelte';
	import LeakFile from './leak_file.svelte';
	import Platforms from './platforms.svelte';
	import ShareDate from './share_date.svelte';
	import SubmitButton from './submit_button.svelte';

	let context: string;
	let leakFilePath: string;
	let shareDateMS: number;
	let leakers: Array<string>;
	let platforms: Array<string>;
</script>

<form>
	<Context value={context} />
	<ShareDate value={shareDateMS} />
	<Platforms value={platforms} />
	<Leakers value={leakers} />
	<LeakFile value={leakFilePath} />

	<SubmitButton />
</form>
