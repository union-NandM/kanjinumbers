import React from "react";
import ModeType from "../type/ModeType";

type Props = {
  mode: ModeType;
  setMode: (mode: ModeType) => void;
};

// 叩くAPIを切り替えるスイッチ
const ModeSwitch = (props: Props) => {
  const handleClick_N2K = () => {
    props.setMode("n2k");
  };
  const handleClick_K2N = () => {
    props.setMode("k2n");
  };

  return (
    <div className="flex container mx-auto border-2 w-fit rounded-md border-blue-500 mb-10">
      <div
        onClick={handleClick_N2K}
        className={`px-3 cursor-pointer select-none ${
          props.mode === "n2k" ? "bg-blue-500 text-white" : ""
        }`}
      >
        ア→漢
      </div>
      <div
        onClick={handleClick_K2N}
        className={`px-3 cursor-pointer select-none ${
          props.mode === "k2n" ? "bg-blue-500 text-white" : ""
        }`}
      >
        漢→ア
      </div>
    </div>
  );
};

export default ModeSwitch;
