<script lang="ts">
	import type { ToBrew } from '../lib/types';
	import { tobrews } from '$lib/tobrews';

	let brews: ToBrew[];

	tobrews.subscribe((b) => {
		brews = b;
	});
</script>

<ul>
	<li>hello</li>
</ul>
