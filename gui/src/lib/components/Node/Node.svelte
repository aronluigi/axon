<script lang="ts">
	import { type NodeProps } from '@xyflow/svelte';
	import { Handle } from '$lib';
	import { PortType, type Node } from '$lib/types';

	type $$Props = NodeProps & Node;
	export let data: $$Props['data'];
	export let id: $$Props['id'];
	// export let selected: $$Props['selected'] = false;
	// export let draggable: $$Props['draggable'] = undefined;
	// export let selectable: $$Props['selectable'] = undefined;
	// export let deletable: $$Props['deletable'] = true;
	// export let dragging: boolean = false;
	// export let type: $$Props['type'] = 'default';
	export let sourcePosition: $$Props['sourcePosition'] = undefined;
	// export let targetPosition: $$Props['targetPosition'] = undefined;
	// export let zIndex: $$Props['zIndex'];
	// export let width: $$Props['width'] = undefined;
	// export let height: $$Props['height'] = undefined;
	// export let dragHandle: $$Props['dragHandle'] = undefined;
	// export let parentId: $$Props['parentId'] = undefined;
</script>

<div
	class="shadow-md rounded-md border-2 border-gray-800 bg-gradient-to-br from-slate-900 to-slate-800"
>
	<div class="flex flex-col">
		<div class="py-1 px-4 text-white bg-gradient-to-r from-red-500 to-orange-500 rounded-t-md">
			<span class="text-lg">{data.label}</span>
			<span class="inline-block align-baseline text-xs text-gray-900 pl-1">{data.package}</span>
		</div>
		<div class="text-white text-sm flex flex-row -mx-3 pt-10 pb-4">
			<div class="flex flex-col gap-y-3">
				{#each data.inPorts as port, k}
					<div class="flex flex-row">
						<Handle id={`${id}-in-${k.toString()}`} portType={PortType.in} />
						<div class="px-2">
							<span>{port.name}</span>
							<span class="text-xs text-gray-400">{port.type}</span>
						</div>
					</div>
				{/each}
			</div>
			<div class="pl-10 flex flex-col gap-y-3">
				{#each data.outPorts as port, k}
					<div class="flex flex-row-reverse">
						<Handle id={`${id}-out-${k.toString()}`} portType={PortType.out} />
						<div class="px-2">
							<span>{port.name}</span>
							<span class="text-xs text-gray-400">{port.type}</span>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>
