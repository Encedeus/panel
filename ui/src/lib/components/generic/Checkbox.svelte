<script lang="ts">
    import { onMount } from "svelte";

    function dec2hex (dec) {
        return dec.toString(16).padStart(2, "0");
    }

    function generateId (len) {
        const arr = new Uint8Array((len || 40) / 2);
        window.crypto.getRandomValues(arr);
        return Array.from(arr, dec2hex).join("");
    }

    let thisInput: HTMLInputElement;
    let thisLabel: HTMLLabelElement;
    onMount(() => {
        const rid = `label-${generateId(10)}`;
        thisInput.id = rid;
        thisLabel.htmlFor = rid;
    });
</script>

<div class="checkbox-wrapper">
    <input bind:this={thisInput} type="checkbox" id="check-23"/>
    <label bind:this={thisLabel} for="check-23" style="--size: 26px">
        <svg viewBox="0,0,50,50">
            <path d="M5 30 L 20 45 L 45 5"></path>
        </svg>
    </label>
</div>

<style lang="postcss">
    .checkbox-wrapper *,
    .checkbox-wrapper *:after,
    .checkbox-wrapper *:before {
        box-sizing: border-box;
    }

    .checkbox-wrapper input {
        position: absolute;
        opacity: 0;
    }

    .checkbox-wrapper input:checked + label svg path {
        stroke-dashoffset: 0;
    }

    .checkbox-wrapper input:checked + label {
        background: #fff;
    }


    .checkbox-wrapper input:focus + label {
        transform: scale(1.03);
    }

    .checkbox-wrapper input + label {
        display: block;
        border: 2px solid #fff;
        width: var(--size);
        height: var(--size);
        border-radius: 6px;
        cursor: pointer;
        transition: all .2s ease;
    }

    .checkbox-wrapper input + label:active {
        transform: scale(1.05);
        border-radius: 12px;
    }

    .checkbox-wrapper input + label svg {
        pointer-events: none;
        padding: 5%;
    }

    .checkbox-wrapper input + label svg path {
        fill: none;
        stroke: #000;
        stroke-width: 4px;
        stroke-linecap: round;
        stroke-linejoin: round;
        stroke-dasharray: 100;
        stroke-dashoffset: 101;
        transition: all 250ms cubic-bezier(1,0,.37,.91);
    }
</style>
