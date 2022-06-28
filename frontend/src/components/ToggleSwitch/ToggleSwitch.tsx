import { useState } from 'react';
import './ToggleSwitch.css';

interface Props {
  value: boolean;
  onChange: (value: boolean) => void;
  isToggled: boolean;
}

export function ToggleSwitch(props: Props) {
  const [isToggled, setIsToggled] = useState(props.value);
  const onToggle = () => {
    setIsToggled(!isToggled);
    props.onChange(!isToggled);
  };
  return (
    <label className="toggle-switch">
      <input type="checkbox" checked={isToggled} onChange={onToggle} />
      <span className="switch" />
    </label>
  );
}
