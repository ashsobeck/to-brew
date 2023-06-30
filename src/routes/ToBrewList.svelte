<script lang="ts">
	import type { ToBrew } from '../lib/types';
	import type { Writable } from 'svelte/store';
	export let store: Writable<ToBrew>;
	export let done: boolean;
</script>

<ul>
	<li>hello</li>
</ul>
