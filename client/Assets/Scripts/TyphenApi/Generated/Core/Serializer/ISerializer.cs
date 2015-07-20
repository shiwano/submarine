namespace TyphenApi
{
    public interface ISerializer
    {
        byte[] Serialize(object type);
        T Deserialize<T>(byte[] bytes) where T : new();
    }
}
