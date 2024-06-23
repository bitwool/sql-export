// DataDisplayComponent.tsx
import React from 'react';
import {main} from '../wailsjs/go/models';

export default function DataDisplayComponent({data }: {data: Array<main.Items>}) {
  return (
    <div>
    {data.map((item) => (
      item.items.map((item) => (
        <div>{item.dbType} - {item.filed} - {item.value}</div>
      ))
    ))}
    </div>
  );
};


