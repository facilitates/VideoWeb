// idl/videoweb.thrift
namespace go videoweb

struct BaseResp {
    1: i64 code;
    2: string msg;
}

struct Auth {
    1: string username;
    2: string password;
}

struct RegisterRequest {
    1: Auth Auth;
}

struct RegisterResponse {
    1: BaseResp base;
}

struct User {
    1: string id;
    2: string username;
    3: string avatar_url;
    4: string created_at;
    5: string updated_at;
    6: string deleted_at;
}

struct LoginRequest {
    1: Auth Auth;
    2: string code;
}

struct LoginResponse {
    1: BaseResp base;
    2: User data;
}

struct InfoRequest {
    1: string user_id;
}

struct InfoResponse {
    1: BaseResp base;
    2: User data;
}

struct UploadAvatarResponse {
    1: BaseResp base;
    2: User data;
}

struct Mfa {
    1: string secret;
    2: string qrcode;
}

struct MfaResponse {
    1: BaseResp base;
    2: Mfa data;
}

struct MfaBindRequest {
    1: string code (api.form="code");
    2: string secret (api.form="secret");
}

struct MfaBindResponse {
    1: BaseResp base (api.body="base");
}

struct VideoFeedRequest {
    1: string latest_time (api.query="latest_time");
}

struct Video{
    1: string Id (api.body="id");
    2: string UserId (api.body="user_id");
    3: string VideoUrl (api.body="video_url");
    4: string CoverUrl (api.body="cover_url");
    5: string Title (api.body="title");
    6: string Description (api.body="description");
    7: string VisitCount (api.body="visit_count");
    8: string LikeCount (api.body="like_count");
    9: string CommentCount (api.body="comment_count");
    10: string CreatedAt (api.body="created_at");
    11: string UpdatedAt (api.body="updated_at");
    12: string DeletedAt (api.body="deleted_at");
}(api.body="items")

struct VideoFeedResponse {
    1: BaseResp base (api.body="base");
    2: list<Video> items (api.body="items");
}
//定义文件
struct UploadVideoRequest {
    1: binary file;
    2: string title;
    3: string description;
}

struct UploadVideoResponse {
    1: BaseResp base;
}

struct VideoListRequest {
    1: string user_id;
    2: i64 page_num;
    3: i64 page_size;
}

struct VideoListResponse {
    1: BaseResp base;
    2: list<Video> data;
    3: i64 total;
}

struct PopularVideoRequest {
    1: i64 page_size;
    2: i64 page_num;
}

struct PopularVideoResponse {
    1: BaseResp base;
    2: list<Video> data;
}

struct SearchVideoRequest {
    1: string keywords;
    2: i64 page_size;
    3: i64 page_num;
    4: i64 from_data;
    5: i64 to_data;
    6: string username;
}

struct SearchVideoResponse {
    1: BaseResp base;
    2: list<Video> data;
    3: i64 total;
}

struct LikeActionRequest {
    1: string video_id;
    2: string comment_id;
    3: string action_type;
}

struct LikeActionResponse {
    1: BaseResp base;
}

struct LikeListRequest {
    1: string user_id;
    2: i64 page_size;
    3: i64 page_num;
}

struct LikeListResponse {
    1: BaseResp base;
    2: list<Video> data;
}

struct CommentPublishRequest {
    1: string video_id;
    2: string comment_id;
    3: string content;
}

struct CommentPublishResponse {
    1: BaseResp base;
}

struct CommentListRequest {
    1: string video_id;
    2: string comment_id;
    3: i64 page_size;
    4: i64 page_num;
}

struct Comment {
    1: string id;
    2: string user_id;
    3: string video_id;
    4: string parent_id;
    5: i64 like_count;
    6: i64 child_count;
    7: string content;
    8: string created_at;
    9: string updated_at;
    10: string deleted_at;
}(api.body="comment")

struct CommentListResponse {
    1: BaseResp base;
    2: list<Comment> data;
}

struct CommentDeleteRequest {
    1: string video_id;
    2: string comment_id;
}

struct CommentDeleteResponse {
    1: BaseResp base;
}

struct RelationActionRequest {
    1: string to_user_id;
    2: i64 action_type;
}

struct RelationActionResponse {
    1: BaseResp base;
}

struct FollowingListRequest {
    1: string user_id;
    2: i64 page_num;
    3: i64 page_size;
}

struct Following {
    1: string id;
    2: string user_name;
    3: string avatar_url;
}

struct FollowingListResponse {
    1: BaseResp base;
    2: Data data;
}

struct FollowerListRequest{
    1: string user_id;
    2: i64 page_num;
    3: i64 page_size;
}

struct FollowerListResponse {
    1: BaseResp base;
    2: Data data;
}

struct FriendsListRequest {
    1: i64 page_num;
    2: i64 page_size;
}

struct Data {
    list<Following> items;
    i64 total;
}

struct FriendsListResponse {
    1: BaseResp base;
    2: Data data;
}

struct ErrResp {
    1: BaseResp base;
}

struct Type1 {
    1: string content;
    2: string id;
}

struct Type2 {
    1: i64 pagenum;
    2: i64 pagesize;
    3: string id;
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="/user/register");
    LoginResponse Login(1: LoginRequest req) (api.post="/user/login");
    InfoResponse Info(1: InfoRequest req) (api.get="/user/info");
    UploadAvatarResponse AvatarUpload() (api.put="/user/avatar/upload")
    MfaResponse MfaQrcode() (api.get="/auth/mfa/qrcode")
    MfaBindResponse MfaBind(1: MfaBindRequest req) (api.post="/auth/mfa/bind")
}

service VideoService {
    VideoFeedResponse VideoFeed(1: VideoFeedRequest req) (api.get="/video/feed")
    UploadVideoResponse UploadVideo(1: UploadVideoRequest req) (api.post="/video/publish")
    VideoListResponse VideoList(1: VideoListRequest req) (api.get="/video/list")
    PopularVideoResponse PopularVideo(1: PopularVideoRequest req) (api.get="/video/popular")
    SearchVideoResponse SearchVideo(1: SearchVideoRequest req) (api.post="/video/search")
}

service ActionService {
    LikeActionResponse LikeAction(1: LikeActionRequest req) (api.post="/like/action")
    LikeListResponse LikeList(1: LikeListRequest req) (api.get="/like/list")
    CommentPublishResponse CommentPublish(1: CommentPublishRequest req) (api.post="/comment/publish")
    CommentListResponse CommentList(1: CommentListRequest req) (api.get="/comment/list")
    CommentDeleteResponse CommentDelete(1: CommentDeleteRequest req) (api.delete="/comment/delete")
}

service SocialService {
    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post="/relation/action")
    FollowingListResponse FollowingList(1: FollowingListRequest req) (api.get="/following/list")
    FollowerListResponse FollowerList(1: FollowerListRequest req) (api.get="/follower/list")
    FriendsListResponse FriendsList(1: FriendsListRequest req) (api.get="/friends/list")
}