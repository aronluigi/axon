<script lang="ts">
	import { useConnection } from '@xyflow/svelte';
	import { getBezierPath } from '@xyflow/system';

	const connection = useConnection();
	let path: string | null = $state(null);

	$effect(() => {
		const { from, to, fromPosition, toPosition } = $connection;
		if (from == null || to == null) return;

		const pathParams = {
			sourceX: from.x,
			sourceY: from.y,
			sourcePosition: fromPosition,
			targetX: to.x,
			targetY: to.y,
			targetPosition: toPosition
		};

		[path] = getBezierPath(pathParams);
	});
</script>

{#if path}
	<path fill="none" stroke-width={4} class="animated stroke-green-700" d={path} />
	<circle
		cx={$connection.to ? $connection.to.x : 0}
		cy={$connection.to ? $connection.to.y : 0}
		r={7}
		stroke-width={3}
		class="stroke-green-700 fill-green-500"
	/>
{/if}
