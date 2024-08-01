<script lang="ts">
	import { nodesStore } from '$lib';
	import { setContext } from 'svelte';
	import { writable, type Writable } from 'svelte/store';
	import { SvelteFlow, Background, Controls, type ColorMode, type Node } from '@xyflow/svelte';

	import '../tailwind.css';
	import '@xyflow/svelte/dist/style.css';

	import { FnNode } from '$lib';

	const edges = writable([]);
	const nodes: Writable<Node[]> = writable([]);
	const colorMode: ColorMode = 'dark';

	const nodeTypes = {
		custom: FnNode
	};

	setContext('nodes', nodes);
</script>

<main class="font-mono">
	<SvelteFlow nodes={nodesStore} {edges} {colorMode} {nodeTypes}>
		<Background />
		<Controls />
	</SvelteFlow>
</main>

<style>
	main {
		height: 100vh;
	}
</style>
