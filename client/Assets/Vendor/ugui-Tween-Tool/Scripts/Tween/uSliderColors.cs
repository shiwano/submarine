//
// Copyright (c) 2015 Tomtc123
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.Events;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Tween/Slider Colors(uTools)")]
	public class uSliderColors : MonoBehaviour {
		public Image target;
		public Color[] colors = new Color[]{Color.red, Color.yellow, Color.green};

		Slider mSlider;

		// Use this for initialization
		void Start () {
			mSlider = GetComponent<Slider>();
			if (mSlider == null) {
				Debug.LogError(" 'uSliderColors' can't find 'Slider'.");
				return;
			}
			if (target == null) {
				target = mSlider.GetComponentInChildren<Image>();
			}
			UnityAction<float> valueChange = new UnityAction<float>(OnValueChanged);
			mSlider.onValueChanged.AddListener(valueChange);
			OnValueChanged(mSlider.value);
		}

		public void OnValueChanged(float value) {
			float val = value * (colors.Length - 1);
			int startIndex = Mathf.FloorToInt(val);
			Color c = colors[0];
			if ( (startIndex + 1) < colors.Length) {
				c = Color.Lerp(colors[startIndex], colors[startIndex+1], val - startIndex);
			}
			else if (startIndex < colors.Length) {
				c = colors[startIndex];
			}
			target.color = c;
		}

	}
}