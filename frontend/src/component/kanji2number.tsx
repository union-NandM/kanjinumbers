import Converter from "./Converter";

// 漢数字からアラビア数字に変換するコンポーネント
const Kanji2Number = () => (
  <Converter
    from="漢数字"
    to="アラビア数字"
    api_uri="kanji2number"
    inputType="text"
  />
);

export default Kanji2Number;
