package http

//// GetComment - retrieve a comment by ID
//func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	i, err := strconv.ParseUint(id, 10, 64)
//	if err != nil {
//		sendErrorResponse(w, "Unable to parse UINT from ID", err)
//		return
//	}
//
//	comment, err := h.Service.GetComment(uint(i))
//	if err != nil {
//		sendErrorResponse(w, "Error Retrieving Comment by ID", err)
//		return
//	}
//
//	if err := sendOkResponse(w, Response{Message: "Successfully Retrieved Comment"}); err != nil {
//		panic(err)
//	}
//	err = json.NewEncoder(w).Encode(comment)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// GetAllComments - Retrieves all comments from the comment service
//func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
//	comments, err := h.Service.GetAllComments()
//	if err != nil {
//		sendErrorResponse(w, "Error to  retrieve all comments", err)
//		return
//	}
//	if err := sendOkResponse(w, Response{Message: "Successfully Retrieved All Comments"}); err != nil {
//		panic(err)
//	}
//	err = json.NewEncoder(w).Encode(comments)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// PostComment - adds a new comment
//func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
//	var comment comment.Comment
//	err := json.NewDecoder(r.Body).Decode(&comment)
//	if err != nil {
//		sendErrorResponse(w, "Failed to decode JSON Body", err)
//		return
//	}
//
//	comment, err = h.Service.PostComment(comment)
//	if err != nil {
//		sendErrorResponse(w, "Failed to post new comment", err)
//		return
//	}
//	if err := sendOkResponse(w, Response{Message: "Successfully Posted Comment"}); err != nil {
//		panic(err)
//	}
//	err = json.NewEncoder(w).Encode(comment)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// UpdateComment - Updates comment by ID
//func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
//	var comment comment.Comment
//	err := json.NewDecoder(r.Body).Decode(&comment)
//	if err != nil {
//		sendErrorResponse(w, "Failed to decode JSON Body", err)
//		return
//	}
//
//	vars := mux.Vars(r)
//	id := vars["id"]
//	commentID, err := strconv.ParseUint(id, 10, 64)
//	if err != nil {
//		sendErrorResponse(w, "Failed to parse uint from ID", err)
//		return
//	}
//
//	comment, err = h.Service.UpdateComment(uint(commentID), comment)
//	if err != nil {
//		sendErrorResponse(w, "Failed to update comment", err)
//		return
//	}
//	if err := sendOkResponse(w, Response{Message: "Successfully Updated Comment"}); err != nil {
//		panic(err)
//	}
//	err = json.NewEncoder(w).Encode(comment)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// DeleteComment - deletes comment by ID
//func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//	commentID, err := strconv.ParseUint(id, 10, 64)
//	if err != nil {
//		sendErrorResponse(w, "Failed to parse uint from ID", err)
//		return
//	}
//
//	if err := h.Service.DeleteComment(uint(commentID)); err != nil {
//		sendErrorResponse(w, "Failed to delete comment by comment ID", err)
//		return
//	}
//	if err := sendOkResponse(w, Response{Message: "Successfully Deleted Comment"}); err != nil {
//		panic(err)
//	}
//	if err := json.NewEncoder(w).Encode(Response{Message: "Comment Successfully Deleted"}); err != nil {
//		panic(err)
//	}
//}