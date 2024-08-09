<script lang="ts">
	import { SvelteFlowProvider } from '@xyflow/svelte';

	import '../tailwind.css';
	import '@xyflow/svelte/dist/style.css';

	import Flow from './Flow.svelte';
</script>

<SvelteFlowProvider>
	<Flow />
</SvelteFlowProvider>
