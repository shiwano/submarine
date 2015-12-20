using UnityEngine;
using System.Collections;
using System.Collections.Generic;

namespace uTools {
	public class uTweenPath : uTweenValue {

		public RectTransform target;
		public List<Vector3> paths;

		int mIndex = -1;
		int mPathsCount = 0;
		bool mCache = false;

		void Cache () {
			mCache = true;
			if (paths.Count > 1) {
				mPathsCount = paths.Count - 1;
			}
			if (target == null) {
				target = GetComponent<RectTransform>();
			}
			from = 0;
			to = mPathsCount;
		}
		
		// Update is called once per frame
		void Update () {
		
		}


		protected override void ValueUpdate (float _factor, bool _isFinished)
		{
			if (!mCache) { Cache();}
			pathIndex = Mathf.FloorToInt(_factor);
			Debug.Log(pathIndex);
		}

		int pathIndex {
			get { return mIndex;}
			set {
				if (mIndex != value) {
					mIndex = value;
					Debug.Log(target.localPosition);
					uTweenPosition.Begin(target.gameObject, target.localPosition, paths[mIndex], duration/paths.Count).loopStyle = LoopStyle.Loop;
				}
			}
		}

	}
}
