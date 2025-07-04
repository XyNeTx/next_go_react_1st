'use client'
import React from "react";

export default function Click() {
    const [count, setCount] = React.useState(1);

    function countClicked(){
        return setCount(count + 1);
    }
    return (
        <button 
            onClick={countClicked} 
            className="bg-black/[.05] dark:bg-white/[.06] px-1 py-0.5 rounded font-[family-name:var(--font-geist-mono)] font-semibold"
        >
            Click me {count}
        </button>
    );
}