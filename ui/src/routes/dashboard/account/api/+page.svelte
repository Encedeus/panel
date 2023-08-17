<script lang="ts">
    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
    import PlusIcon from "$lib/components/heroicons/PlusIcon.svelte";
</script>

<main class="p-6">
    <div class="flex flex-row gap-8 items-end">
        <div>
            <CardHeader size="lg" className="mb-4">
                API Credentials
            </CardHeader>
            <Card width="lg" height="sm">
                <span slot="title" class="text-sm">
                    Create API Key
                </span>
                <PlusIcon slot="icon"/>
            </Card>
        </div>        
        <div>
            <Card width="lg" height="sm">
                <span slot="title" class="text-sm">
                    API Keys
                </span>
                <KeyIcon slot="icon"/>
            </Card>
        </div>
    </div>
</main>
