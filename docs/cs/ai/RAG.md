# RAG

检索增强生成(Retrieval-Augmented Generation)

先从资料库检索相关内容，在基于这些内容生成答案

## 流程

提问前: 分片 + 索引

提问后: 召回 + 重排 + 生成

+ 分片(对数据进行切割，分成不同的片段)  字数，段落，页码，章节
+ 索引(通过Embeddding将片段转化为向量，把向量和文本存储向量数据库)
+ 召回(计算向量相似度，余弦相似度，欧式距离，点积，取排名靠前的)
+ 重排(用cross-encoder从上面挑选部分分片)
+ 生成(发给大模型进行推理)

## 用途

+ 智能客服
+ 知识助手

## FAQ

1. 不使用向量数据库也可以做

OpenAI 方案: AgentRAG, 使用多个模型

基于大模型对所有数据的理解，然后进行多次切割，内容挑选，生成答案，然后验证

## demo

```py
from typing import List

def split_into_chunks(doc_file: str) -> List[str]:
    with open(doc_file, 'r') as file:
        content = file.read()
    return [chunk for chunk in content.split("\n\n")]

chunks = split_into_chunks("doc.md")

for i, chunk in enumerate(chunks):
    print(f"[{i}] {chunk}\n")

from sentence_transformers import SentenceTransformer

embedding_model = SentenceTransformer("shibing624/text2vec-base-chinese")

def embed_chunk(chunk: str) -> List[float]:
    embedding = embedding_model.encode(chunk)
    return embedding.tolist()

embeddings = [embed_chunk(chunk) for chunk in chunks]


import chromadb

chromadb_client = chromadb.EphemeralClient()
chromadb_collection = chromadb_client.get_or_create_collection(name="default")

def save_embeddings(chunks: List[str], embeddings: List[List[float]]) -> None:
    ids = [str(i) for i in range(len(chunks))]
    chromadb_collection.add(
        documents = chunks,
        embeddings = embeddings,
        ids = ids
    )

save_embeddings(chunks, embeddings)

def retrieve(query: str, top_k: int) -> List[str]:
    query_embedding = embed_chunk(query)
    results = chromadb_collection.query(
        query_embeddings = [query_embedding],
        n_results = top_k
    )
    return results['documents'][0]

query = "哆啦A梦"

retrieved_chunks = retrieve(query, 5)
for i, chunk in enumerate(retrieved_chunks):
    print(f"[{i}] {chunk}\n")

from sentence_transformers import CrossEncoder

def rerank(query: str, retrieved_chunks: List[str], top_k: int) -> List[str]:
    cross_encoder = CrossEncoder('cross-encoder/mmarco-mMiniLMv2-L12-H384-v1')
    pairs = [(query, chunk) for chunk in retrieved_chunks]
    scores = cross_encoder.predict(pairs)

    scored_chunks = list(zip(retrieved_chunks, scores))
    scored_chunks.sort(key=lambda pair: pair[1], reverse=True)
    return [chunk for chunk, _ in scored_chunks][:top_k]

reranked_chunks = rerank(query, retrieved_chunks, 3)

for i, chunk in enumerate(reranked_chunks):
    print(f"{[i]} {chunk}\n")


from dotenv import load_dotenv
from google import genai

load_dotenv()
google_client = genai.Client()

def generate(query: str, chunks: List[str]) -> str:
    prompt = f"""你是一位知识助手, 请根据用户的问题和下列片段生成准确的回答。
    用户问题: {query}

    相关片段:
    {"\n\n".join(chunks)}
    请基于以上内容作答，不要编造信息"""

    print(f"{prompt}\n\n---\n")

    response = google_client.models.generate_content(
        model="gemini-2.5-flash",
        contents=prompt
    )

    return response.text

answer = generate(query, reranked_chunks)
print(answer)

```