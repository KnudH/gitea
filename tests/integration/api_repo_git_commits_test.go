// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
	session.MakeRequest(t, req, http.StatusNotFound)
	session.MakeRequest(t, req, http.StatusUnprocessableEntity)
	session.MakeRequest(t, req, http.StatusNotFound)
		session.MakeRequest(t, req, http.StatusOK)
	resp := session.MakeRequest(t, req, http.StatusOK)
	resp := session.MakeRequest(t, req, http.StatusOK)
	resp := session.MakeRequest(t, req, http.StatusOK)
	resp := session.MakeRequest(t, reqDiff, http.StatusOK)
	resp = session.MakeRequest(t, reqPatch, http.StatusOK)
	resp := session.MakeRequest(t, req, http.StatusOK)