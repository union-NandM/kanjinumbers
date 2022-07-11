import Converter from "./Converter";

// アラビア数字から漢数字に変換するコンポーネント
const Number2Kanji = () => (
  <Converter
    from="アラビア数字"
    to="漢数字"
    api_uri="number2kanji"
    inputType="number"
  />
);

export default Number2Kanji;
