<script lang="ts">
	import { beans } from '$lib/stores.js';
	import type { Bean } from '$lib/types.js';
	import type { ActionData } from '../$types.js';
	import AddBean from './AddBean.svelte';
	import BeanList from './BeanList.svelte';

	type BeanData = {
		beans: Bean[];
	};

	export let data: BeanData;
	// export let form: ActionData;

	console.log('getting beans...');
	console.log(data.beans);
	beans.set(data.beans);
</script>

<div class="w-full h-screen grid gap-x-1 items-center p-2">
	<div class="bg-surface-500/30 h-full card p-4">
		Beans
		<BeanList />
		<AddBean />
	</div>
	<!-- <div class="bg-surface-500/30 h-full card p-4">current bean</div> -->
</div>
