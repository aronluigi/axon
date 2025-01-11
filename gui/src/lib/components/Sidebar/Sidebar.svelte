<script lang="ts">
	import { type NodeObj } from '$lib/types';

	interface Props {
		nodes: NodeObj[];
		onDragStart(e: DragEvent, uuid: string): void;
	}

	let { nodes, onDragStart }: Props = $props();
</script>

<aside class="sidebar text-sm border-r-gray-800 border-r-4 text-white p-2">
	<div
		class="p-2.5 flex items-center rounded-md px-4 duration-300 cursor-pointer bg-gray-700 text-white"
	>
		<i class="bi bi-search text-sm"></i>
		<input
			type="text"
			placeholder="Search"
			class="text-sm ml-4 w-full bg-transparent focus:outline-none"
		/>
	</div>
	{#each nodes as node, i}
		<div
			role="menuitem"
			tabindex={i}
			draggable={true}
			ondragstart={(event) => onDragStart(event, node.uuid)}
			class="p-2.5 mt-3 flex items-center rounded-md px-4 duration-300 bg-gray-900 cursor-pointer hover:bg-gray-800"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="size-5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m21 7.5-9-5.25L3 7.5m18 0-9 5.25m9-5.25v9l-9 5.25M3 7.5l9 5.25M3 7.5v9l9 5.25m0-9v9"
				/>
			</svg>
			<div class="ml-4">
				<p class="text-sm">{node.label}</p>
				<p class="text-xs text-gray-500">{node.package}</p>
			</div>
		</div>
	{/each}
</aside>

<style>
	.sidebar {
		background-color: #141414;
	}
</style>
