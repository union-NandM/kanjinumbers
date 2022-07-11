import { ChangeEvent, useState } from "react";

type Props = {
  from: "アラビア数字" | "漢数字";
  to: "アラビア数字" | "漢数字";
  inputType: "number" | "text";
  apiUri: "number2kanji" | "kanji2number";
};

const Converter = (props: Props) => {
  // 変換元数値
  const [fromNumber, setFromNumber] = useState<string>("");
  // 変換先数値
  const [toNumber, setToNumber] = useState<string>("");

  // 入力
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFromNumber(e.target.value);
  };

  // fetch
  const handleClick = async () => {
    if (fromNumber === "") {
      return;
    }
    const res = await fetch(
      encodeURI(
        `https://rxyfko3ctb.execute-api.ap-northeast-1.amazonaws.com/v1/${props.apiUri}/${fromNumber}`
      ),
      {
        mode: "cors",
      }
    );

    if (!res.ok) {
      alert("エラーが発生しました");
    } else if (res?.status === 204) {
      alert("変換できませんでした。正しい形式で入力してください。");
    } else {
      const data = await res.json();
      setToNumber(data.data);
    }
  };

  return (
    <>
      <div>
        <div className="pb-3 select-none">{props.from}</div>
        <input
          type={props.inputType}
          onChange={handleChange}
          className="border-b-2 border-gray-200 outline-none focus:border-gray-600 w-full"
          placeholder={`0〜9,999,999,999,999,999の${props.from}`}
        />
      </div>
      <div className="w-min mx-auto my-10">⬇️</div>
      <div>
        <div className="pb-3 select-none">{props.to}</div>
        <div className="min-h-[6rem] text-xl leading-8">
          {/* 数値のときはカンマ区切りで表示（Number.MAX_SAFE_INTEGERを超えることがあるため正規表現で実装） */}
          {toNumber.replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,")}
        </div>
      </div>
      <div className="flex justify-center">
        <input
          type="button"
          value="変換"
          onClick={handleClick}
          className="bg-blue-500 text-white rounded-md py-1 px-2 select-none"
        />
      </div>
    </>
  );
};

export default Converter;
