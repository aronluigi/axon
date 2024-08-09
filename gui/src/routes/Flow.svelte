<script lang="ts">
	import { nodesStore, availableNodes, FnNode, Sidebar, type NodeObj } from '$lib';
	import { writable } from 'svelte/store';
	import { SvelteFlow, Background, Controls, type ColorMode, useSvelteFlow } from '@xyflow/svelte';
	import { onNodeDrop, onNodeDragOver, onNodeDragStart } from './NodeEvents';
	import { onMount } from 'svelte';

	let availableNodesList: NodeObj[] = [];
	availableNodes.subscribe((v) => {
		availableNodesList = Object.keys(v).map((k) => v[k]);
	});

	const edges = writable([]);
	const colorMode: ColorMode = 'dark';

	const nodeTypes = {
		custom: FnNode
	};

	const { screenToFlowPosition } = useSvelteFlow();
</script>

<main class="font-mono flex flex-row">
	<SvelteFlow
		nodes={nodesStore}
		on:dragover={onNodeDragOver}
		on:drop={onNodeDrop(screenToFlowPosition)}
		{edges}
		{colorMode}
		{nodeTypes}
	>
		<Background />
		<Controls />
	</SvelteFlow>
	<Sidebar nodes={availableNodesList} onDragStart={onNodeDragStart} />
</main>

<style>
	main {
		height: 100vh;
	}
</style>
