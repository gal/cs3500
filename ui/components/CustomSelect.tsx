import {Tag} from 'api/types';
import {FieldProps} from 'formik';
import React from 'react';
import Select from 'react-select';

interface Option {
  label: string;
  value: string;
  selected?: boolean;
}

interface CustomSelectProps extends FieldProps {
  className?: string;
  options: Option[];
  alreadySelectedOptions?: Tag[];
  isMulti?: boolean;
  placeholder?: string;
  label: string;
  muted?: boolean;
  description?: string;
}

export const CustomSelect = ({
  className,
  placeholder,
  field,
  form,
  options,
  label,
  muted,
  description,
  isMulti = false,
}: CustomSelectProps) => {
  React.useEffect(() => {
    const selectedOptions = options.filter(
        (option) => option.selected === true,
    );
    if (selectedOptions) {
      form.setFieldValue(
          field.name,
          selectedOptions?.map((option: Option) => option.value),
      );
    }
  }, []);

  const onChange = (option: Option | Option[]) => {
    form.setFieldValue(
        field.name,
      isMulti ?
        (option as Option[]).map((item: Option) => item.value) :
        (option as Option).value,
    );
  };

  const getValue = () => {
    if (options) {
      return isMulti ?
        options.filter((option) => field.value.indexOf(option.value) >= 0) :
        options.find((option) => option.value === field.value);
    } else {
      return isMulti ? [] : ('' as any);
    }
  };

  return (
    <Select
      name={field.name}
      value={getValue()}
      onChange={onChange}
      placeholder={placeholder}
      options={options}
      isMulti={isMulti}
    />
  );
};

export default CustomSelect;
