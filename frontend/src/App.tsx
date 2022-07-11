import React, { useState } from "react";
import ModeSwitch from "./component/ModeSwitch";
import Kanji2Number from "./component/Kanji2Number";
import Number2Kanji from "./component/Number2Kanji";
import ModeType from "./type/ModeType";

const App = () => {
  const [mode, setMode] = useState<ModeType>("n2k");

  return (
    <>
      <header className="w-full bg-blue-200 text-lg mb-10 font-title">
        kanjinumbers.com
      </header>
      <div className="w-full sm:w-80 sm:w-fit-sm m-auto px-3 font-ja">
        <ModeSwitch
          mode={mode}
          setMode={(mode: ModeType) => {
            setMode(mode);
          }}
        />
        {mode === "n2k" ? <Number2Kanji /> : <Kanji2Number />}
      </div>
    </>
  );
};

export default App;
