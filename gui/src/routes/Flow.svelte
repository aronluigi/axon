<script lang="ts">
	import {
		nodesStore,
		availableNodes,
		FnNode,
		ConnectionLine,
		Sidebar,
		type NodeObj,
		getAvailableNodes
	} from '$lib';
	import { fromStore, writable } from 'svelte/store';
	import { SvelteFlow, Background, Controls, type ColorMode, useSvelteFlow } from '@xyflow/svelte';
	import { onNodeDrop, onNodeDragOver, onNodeDragStart } from './NodeEvents';
	import { onMount } from 'svelte';

	let sidebarNodes: NodeObj[] = $derived.by(() => {
		const res: NodeObj[] = [];
		const vals = fromStore(availableNodes);

		for (let k in vals.current) res.push(vals.current[k]);

		return res;
	});

	const { screenToFlowPosition } = useSvelteFlow();
	const edges = writable([]);
	const colorMode: ColorMode = 'dark';
	const nodeTypes = {
		custom: FnNode
	};

	onMount(async () => availableNodes.set(await getAvailableNodes()));
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
		<Controls />
		<ConnectionLine slot="connectionLine" />
		<Background />
	</SvelteFlow>
	<Sidebar nodes={sidebarNodes} onDragStart={onNodeDragStart} />
</main>

<style>
	main {
		width: 100vw;
		height: 100vh;
	}
</style>
